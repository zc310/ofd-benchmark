#!/usr/bin/env node

import fs from 'fs';
import { createRequire } from 'node:module';

const require2 = createRequire(import.meta.url);
const { convert: convertCjs } = require2('@miconvert/ofd-to-pdf');
const { convert: convertEsm } = await import('@miconvert/ofd-to-pdf');

const args = process.argv.slice(2);
if (args.length < 2) {
    console.log('Usage: ofd2pdf.mjs <input.ofd> <output.pdf>');
    process.exit(1);
}

const inputPath = args[0];
const outputPath = args[1];

try {
    const ofdBytes = fs.readFileSync(inputPath);
    let pdfBytes;
    try {
        pdfBytes = await convertCjs(ofdBytes, { silent: true });
    } catch (cjsError) {
        pdfBytes = await convertEsm(ofdBytes, { silent: true });
    }
    fs.writeFileSync(outputPath, pdfBytes);
    console.log(`Converted: ${inputPath} -> ${outputPath}`);
} catch (error) {
    console.error(`FAILED: ${error.message}`);
    process.exit(1);
}
