#!/usr/bin/env node

const https = require('https');
const fs = require('fs');
const path = require('path');

const GITHUB_REPO = 'fdorantesm/aka';
const BIN_DIR = path.join(__dirname, '..', 'bin');
const BINARY_NAME = 'aka';

// Ensure bin directory exists
if (!fs.existsSync(BIN_DIR)) {
  fs.mkdirSync(BIN_DIR, { recursive: true });
}

function getPlatformInfo() {
  const platform = process.platform;
  const arch = process.arch;

  let binaryName;
  
  if (platform === 'darwin') {
    binaryName = 'aka-darwin';
  } else if (platform === 'linux') {
    binaryName = 'aka-linux';
  } else {
    throw new Error(`Unsupported platform: ${platform} ${arch}`);
  }

  return { binaryName };
}

function getLatestRelease() {
  return new Promise((resolve, reject) => {
    const options = {
      hostname: 'api.github.com',
      path: `/repos/${GITHUB_REPO}/releases/latest`,
      headers: {
        'User-Agent': 'aka-npm-installer'
      }
    };

    https.get(options, (res) => {
      let data = '';

      res.on('data', (chunk) => {
        data += chunk;
      });

      res.on('end', () => {
        try {
          const release = JSON.parse(data);
          resolve(release);
        } catch (error) {
          reject(new Error(`Failed to parse release data: ${error.message}`));
        }
      });
    }).on('error', (error) => {
      reject(new Error(`Failed to fetch latest release: ${error.message}`));
    });
  });
}

function downloadFile(url, destination) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(destination);
    
    https.get(url, {
      headers: {
        'User-Agent': 'aka-npm-installer'
      }
    }, (response) => {
      if (response.statusCode === 302 || response.statusCode === 301) {
        // Follow redirect
        downloadFile(response.headers.location, destination)
          .then(resolve)
          .catch(reject);
        return;
      }

      if (response.statusCode !== 200) {
        reject(new Error(`Failed to download: ${response.statusCode}`));
        return;
      }

      response.pipe(file);

      file.on('finish', () => {
        file.close();
        resolve();
      });
    }).on('error', (error) => {
      fs.unlink(destination, () => {});
      reject(new Error(`Download failed: ${error.message}`));
    });

    file.on('error', (error) => {
      fs.unlink(destination, () => {});
      reject(new Error(`File write failed: ${error.message}`));
    });
  });
}

async function install() {
  try {
    console.log('📦 Installing aka...');
    
    const { binaryName } = getPlatformInfo();
    console.log(`🔍 Detected platform binary: ${binaryName}`);

    console.log('🌐 Fetching latest release...');
    const release = await getLatestRelease();
    const version = release.tag_name;
    console.log(`✨ Latest version: ${version}`);

    const asset = release.assets.find(asset => asset.name === binaryName);
    if (!asset) {
      throw new Error(`Binary ${binaryName} not found in release ${version}`);
    }

    const downloadUrl = asset.browser_download_url;
    const binaryPath = path.join(BIN_DIR, BINARY_NAME);

    console.log(`⬇️  Downloading from ${downloadUrl}...`);
    await downloadFile(downloadUrl, binaryPath);

    // Make binary executable
    fs.chmodSync(binaryPath, '755');

    console.log('✅ aka installed successfully!');
    console.log('');
    console.log('Run "aka --help" to get started.');
  } catch (error) {
    console.error('❌ Installation failed:', error.message);
    process.exit(1);
  }
}

install();
