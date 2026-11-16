#!/usr/bin/env python3
"""Compare converted PDFs against the reference PDF (数科版式 reader output).

Usage:
    python3 tools/compare_pdf.py [dpi] [file...]

Renders each page of both PDFs and computes pixel-level similarity.
"""
import sys
import os
import numpy as np
import fitz
from benchmark_results import OFD_CONVERTER_ROWS, append, write_updates

def render(doc, page_idx, dpi):
    page = doc[page_idx]
    zoom = dpi / 72.0
    fitz.TOOLS.mupdf_warnings(reset=True)
    pix = page.get_pixmap(matrix=fitz.Matrix(zoom, zoom), colorspace=fitz.csRGB, alpha=False)
    warnings = fitz.TOOLS.mupdf_warnings(reset=True)
    if warnings:
        raise ValueError(f'incomplete render on page {page_idx + 1}: {warnings.strip()}')
    img = np.frombuffer(pix.samples, dtype=np.uint8).reshape(pix.height, pix.width, pix.n)
    return img

# intro.pdf（数科导出的参考 PDF）第 14–19 页背景图丢失，像素对比跳过这些页
EXCLUDE_PAGES = {'intro': {14, 15, 16, 17, 18, 19}}
# Repeated content: sample the first three pages for all pixel/foreground metrics.
PAGE_LIMITS = {'1000-pages': 3}

WHITE_THRESHOLD = 250
PAGE_SIZE_TOLERANCE_PT = 0.5


def align_images(ref, src):
    """Pad small raster rounding differences; never crop away document content."""
    if ref.ndim != 3 or src.ndim != 3 or ref.shape[2] != 3 or src.shape[2] != 3:
        raise ValueError('expected RGB images')
    if ref.shape == src.shape:
        return ref, src
    if abs(ref.shape[0] - src.shape[0]) > 1 or abs(ref.shape[1] - src.shape[1]) > 1:
        raise ValueError(f'raster size mismatch: reference={ref.shape}, output={src.shape}')
    h, w = max(ref.shape[0], src.shape[0]), max(ref.shape[1], src.shape[1])
    def pad(image):
        return np.pad(image, ((0, h - image.shape[0]), (0, w - image.shape[1]), (0, 0)),
                      constant_values=255)
    return pad(ref), pad(src)


def foreground(image):
    return np.any(image < WHITE_THRESHOLD, axis=2)


def dilate(mask, radius):
    padded = np.pad(mask, radius)
    expanded = np.zeros_like(mask)
    for dy in range(2 * radius + 1):
        for dx in range(2 * radius + 1):
            expanded |= padded[dy:dy + mask.shape[0], dx:dx + mask.shape[1]]
    return expanded

def similarity(ref, src):
    ref, src = align_images(ref, src)
    diff = np.abs(ref.astype(np.int16) - src.astype(np.int16)).sum(axis=2) / (3 * 255.0)
    mean_diff = diff.mean()
    identical = (diff < 0.01).mean()
    return mean_diff, identical

def content_similarity(ref, src):
    """Reference-normalized color error, including extra output ink.

    The denominator depends only on reference ink (not the union). Errors on
    reference ink and extra output ink are summed separately; output background
    cannot enlarge the denominator. None means both pages have no foreground.
    """
    ref, src = align_images(ref, src)
    ref_mask, src_mask = foreground(ref), foreground(src)
    if not ref_mask.any():
        return 0.0 if src_mask.any() else None
    diff = np.abs(ref.astype(np.int16) - src.astype(np.int16)).sum(axis=2) / (3 * 255.0)
    missing_error = diff[ref_mask].sum()
    extra_error = diff[src_mask & ~ref_mask].sum()
    return max(0.0, 1 - (missing_error + extra_error) / np.count_nonzero(ref_mask)) * 100


def foreground_f1(ref, src, dpi):
    """Tolerant foreground coverage; color error is deliberately not shifted."""
    ref, src = align_images(ref, src)
    ref_mask, src_mask = foreground(ref), foreground(src)
    if not ref_mask.any() and not src_mask.any():
        return None
    if not ref_mask.any() or not src_mask.any():
        return 0.0
    # One pixel at 96 DPI (~0.265 mm); keep the physical tolerance at other DPIs.
    radius = max(1, round(dpi / 96))
    precision = np.count_nonzero(src_mask & dilate(ref_mask, radius)) / np.count_nonzero(src_mask)
    recall = np.count_nonzero(ref_mask & dilate(src_mask, radius)) / np.count_nonzero(ref_mask)
    return 200 * precision * recall / (precision + recall) if precision + recall else 0.0

def compare(ref_path, src_path, dpi):
    if dpi <= 0:
        raise ValueError('DPI must be positive')
    name = os.path.splitext(os.path.basename(ref_path))[0]
    skip = EXCLUDE_PAGES.get(name, set())
    total_diff = 0.0
    total_ident = 0.0
    total_content = 0.0
    total_f1 = 0.0
    content_count = 0
    count = 0
    with fitz.open(ref_path) as ref_doc, fitz.open(src_path) as src_doc:
        if ref_doc.page_count != src_doc.page_count:
            raise ValueError(f'page count mismatch: reference={ref_doc.page_count}, output={src_doc.page_count}')
        page_limit = min(ref_doc.page_count, PAGE_LIMITS.get(name, ref_doc.page_count))
        for i in range(page_limit):
            if i + 1 in skip:
                continue
            ref_rect, src_rect = ref_doc[i].rect, src_doc[i].rect
            if (abs(ref_rect.width - src_rect.width) > PAGE_SIZE_TOLERANCE_PT or
                    abs(ref_rect.height - src_rect.height) > PAGE_SIZE_TOLERANCE_PT):
                raise ValueError(f'page size mismatch on page {i + 1}: reference={ref_rect}, output={src_rect}')
            ref_img = render(ref_doc, i, dpi)
            src_img = render(src_doc, i, dpi)
            # A nearly empty PDF can score deceptively well because most pages are white.
            # Fail pages that lost virtually all foreground content instead of reporting a
            # misleadingly high whole-page similarity.
            ref_ink = np.count_nonzero(np.any(ref_img < 245, axis=2))
            src_ink = np.count_nonzero(np.any(src_img < 245, axis=2))
            if ref_ink >= 16 and src_ink < max(4, ref_ink * 0.01):
                raise ValueError(f'output page {i + 1} is nearly blank compared with reference')
            md, ident = similarity(ref_img, src_img)
            content = content_similarity(ref_img, src_img)
            if content is not None:
                total_content += content
                total_f1 += foreground_f1(ref_img, src_img, dpi)
                content_count += 1
            total_diff += md
            total_ident += ident
            count += 1

    if not count:
        raise ValueError('no comparable pages')
    mean_diff = total_diff / count
    pct = (1 - mean_diff) * 100
    ident = total_ident / count * 100
    # An all-blank document has no content score, not a perfect content score.
    return (pct, mean_diff, ident,
            total_content / content_count if content_count else None,
            total_f1 / content_count if content_count else None)

def main():
    dpi = 96
    args = sys.argv[1:]
    if args and args[0].isdigit():
        dpi = int(args[0])
        args = args[1:]

    files = args or ['hello', 'ano', '999', '1000-pages', 'intro', 'zsbk']
    base = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    odir = os.path.join(base, 'bin', 'output')

    print(f"{'file':<12} {'convert':<16} {'ref_size':>10} {'pdf_size':>10} {'ratio':>6} {'full-page':>10} {'content':>10} {'fg-F1':>10} {'identical':>9}")
    print("-" * 104)
    updates = []
    for name in files:
        ref = os.path.join(base, 'testdata', 'suwell', f'{name}.pdf')
        ref_size = os.path.getsize(ref)
        if name in PAGE_LIMITS:
            print(f'{name}: pixel/foreground metrics sample first {PAGE_LIMITS[name]} pages (total page count still checked)')
        for pdf in sorted(os.listdir(odir)):
            if not pdf.startswith(name + '_') or not pdf.endswith('.pdf'):
                continue
            conv = pdf[len(name)+1:-4]
            src = os.path.join(odir, pdf)
            src_size = os.path.getsize(src)
            ratio = src_size / ref_size if ref_size else 0
            pixel = 'FAILED'
            content_pixel = 'FAILED'
            f1_pixel = 'FAILED'
            try:
                sim, md, ident, content, f1 = compare(ref, src, dpi)
                pixel = f'{sim:.2f}%'
                content_pixel = f'{content:.2f}%' if content is not None else 'FAILED'
                f1_pixel = f'{f1:.2f}%' if f1 is not None else 'FAILED'
                print(f"{name:<12} {conv:<16} {ref_size:>10} {src_size:>10} {ratio:>6.2f} {sim:>9.2f}% {content_pixel:>10} {f1_pixel:>10} {ident:>8.2f}%")
            except Exception as e:
                print(f"{name:<12} {conv:<16} FAILED: {e}")
            row = OFD_CONVERTER_ROWS.get(conv)
            if row:
                append(updates, 'ofd-to-pdf', 'ofd-size-ratio', f'{name}.ofd', row, f'{ratio:.2f}')
                append(updates, 'ofd-to-pdf', 'ofd-pixel-similarity', f'{name}.ofd', row, pixel)
                append(updates, 'ofd-to-pdf', 'ofd-content-similarity', f'{name}.ofd', row, content_pixel)
                append(updates, 'ofd-to-pdf', 'ofd-foreground-f1', f'{name}.ofd', row, f1_pixel)
        print()
    write_updates(base, updates)

if __name__ == '__main__':
    main()
