import { createRequire } from 'node:module';
import { readFileSync, writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';

// pdfjs-dist 通过 process.getBuiltinModule 走 Node fs 读 CMap；Node <22 无此接口
process.getBuiltinModule ??= createRequire(import.meta.url);

const { pdfToOfd } = await import('@hufe921/jsofd/pdf');

const [inPdf, outOfd] = process.argv.slice(2);
if (!inPdf || !outOfd) {
  console.error('Usage: pdf2ofd_jsofd.mjs <in.pdf> <out.ofd>');
  process.exit(1);
}

const here = dirname(fileURLToPath(import.meta.url));
const pdfjs = join(here, 'node_modules', 'pdfjs-dist');

const bytes = new Uint8Array(readFileSync(inPdf));
const doc = await pdfToOfd(bytes, {
  cMapUrl: join(pdfjs, 'cmaps') + '/',
  cMapPacked: true,
  standardFontDataUrl: join(pdfjs, 'standard_fonts') + '/',
});
writeFileSync(outOfd, Buffer.from(doc.output('arraybuffer')));
console.log(`Converted: ${inPdf} -> ${outOfd}`);