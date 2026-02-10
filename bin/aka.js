#!/usr/bin/env node

const { spawn } = require("child_process");
const path = require("path");
const fs = require("fs");

const binaryPath = path.join(__dirname, "aka");

// Check if binary exists
if (!fs.existsSync(binaryPath)) {
  console.error("❌ The aka binary was not installed correctly.");
  console.error("Please try reinstalling: npm install -g aka-cli");
  process.exit(1);
}

// Execute the binary
const child = spawn(binaryPath, process.argv.slice(2), {
  stdio: "inherit",
});

child.on("exit", (code) => {
  process.exit(code || 0);
});

child.on("error", (err) => {
  console.error("❌ Failed to execute aka:", err.message);
  process.exit(1);
});
