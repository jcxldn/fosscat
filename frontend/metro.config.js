// Learn more https://docs.expo.io/guides/customizing-metro
const { getDefaultConfig } = require('expo/metro-config');

/** @type {import('expo/metro-config').MetroConfig} */
const config = getDefaultConfig(__dirname, {
  // [Web-only]: Enables CSS support in Metro.
  isCSSEnabled: true,
});

// Allow .mjs files to work for graphql package
// https://github.com/pocketbase/js-sdk/issues/47#issuecomment-1277479367
config.resolver.sourceExts.push('mjs');

module.exports = config;