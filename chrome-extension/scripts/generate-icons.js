const sharp = require('sharp');
const fs = require('fs');
const path = require('path');

const sizes = [16, 48, 128];
const inputFile = path.join(__dirname, '../public/icon.svg');
const outputDir = path.join(__dirname, '../build');

// 确保输出目录存在
if (!fs.existsSync(outputDir)) {
  fs.mkdirSync(outputDir, { recursive: true });
}

// 为每个尺寸生成图标
sizes.forEach(size => {
  sharp(inputFile)
    .resize(size, size)
    .png()
    .toFile(path.join(outputDir, `icon${size}.png`))
    .then(() => console.log(`Generated icon${size}.png`))
    .catch(err => console.error(`Error generating icon${size}.png:`, err));
}); 