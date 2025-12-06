# Favicon Files

This directory contains favicon files for the Ping application.

## Current Status

The favicon files in this directory need to be replaced with Ping-branded versions. Currently, they may contain Mattermost branding.

## Required Files

The following favicon files are used by the application:

### Default Favicons (used when no unreads/mentions)
- `favicon-default-16x16.png`
- `favicon-default-24x24.png`
- `favicon-default-32x32.png`
- `favicon-default-64x64.png`
- `favicon-default-96x96.png`

### Unread Favicons (used when there are unread messages)
- `favicon-unread-16x16.png`
- `favicon-unread-24x24.png`
- `favicon-unread-32x32.png`
- `favicon-unread-64x64.png`
- `favicon-unread-96x96.png`

### Mention Favicons (used when there are mentions)
- `favicon-mentions-16x16.png`
- `favicon-mentions-24x24.png`
- `favicon-mentions-32x32.png`
- `favicon-mentions-64x64.png`
- `favicon-mentions-96x96.png`

### Other Favicons
- `favicon-16x16.png`
- `favicon-32x32.png`
- `favicon-96x96.png`
- `android-chrome-192x192.png`
- `apple-touch-icon-*.png` (various sizes)

## How to Replace

1. Generate Ping-branded favicons from the Ping logo (`../ping-logo.png`)
2. Replace all the PNG files in this directory with Ping-branded versions
3. Ensure all sizes match the filenames exactly
4. For unread/mention variants, consider adding visual indicators (e.g., red dot, badge)

## Tools

You can use online favicon generators or image editing tools to create favicons from the Ping logo:
- https://realfavicongenerator.net/
- ImageMagick or similar tools for batch conversion

