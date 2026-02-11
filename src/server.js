'use strict';

const http = require('node:http');
const fs = require('node:fs/promises');
const path = require('node:path');

const HOST = process.env.HOST || '127.0.0.1';
const PORT = Number(process.env.PORT || 3000);

const projectRoot = path.resolve(__dirname, '..');
const templatePath = path.join(projectRoot, 'templates', 'index.html');
const publicDir = path.join(projectRoot, 'public');

const MIME_TYPES = {
  '.css': 'text/css; charset=utf-8',
  '.js': 'application/javascript; charset=utf-8',
  '.svg': 'image/svg+xml',
  '.png': 'image/png',
  '.jpg': 'image/jpeg',
  '.jpeg': 'image/jpeg',
  '.ico': 'image/x-icon',
  '.webp': 'image/webp'
};

function setBaseHeaders(res) {
  res.setHeader('X-Content-Type-Options', 'nosniff');
  res.setHeader('Referrer-Policy', 'no-referrer');
  res.setHeader('X-Frame-Options', 'DENY');
}

function sendText(res, statusCode, text, contentType = 'text/plain; charset=utf-8') {
  res.statusCode = statusCode;
  setBaseHeaders(res);
  res.setHeader('Content-Type', contentType);
  res.end(text);
}

async function serveRoot(res) {
  try {
    const html = await fs.readFile(templatePath, 'utf8');
    res.statusCode = 200;
    setBaseHeaders(res);
    res.setHeader('Content-Type', 'text/html; charset=utf-8');
    res.setHeader('Cache-Control', 'no-store');
    res.end(html);
  } catch {
    sendText(res, 500, 'Template not found');
  }
}

async function serveStatic(reqPath, res) {
  const relativePath = reqPath.replace(/^\/static\//, '');
  const normalizedPath = path.normalize(relativePath).replace(/^([.][.][/\\])+/, '');
  const filePath = path.join(publicDir, normalizedPath);

  if (!filePath.startsWith(publicDir)) {
    sendText(res, 403, 'Forbidden');
    return;
  }

  const extension = path.extname(filePath).toLowerCase();
  const contentType = MIME_TYPES[extension];

  if (!contentType) {
    sendText(res, 404, 'Not found');
    return;
  }

  try {
    const file = await fs.readFile(filePath);
    res.statusCode = 200;
    setBaseHeaders(res);
    res.setHeader('Content-Type', contentType);
    res.setHeader('Cache-Control', 'public, max-age=3600');
    res.end(file);
  } catch {
    sendText(res, 404, 'Not found');
  }
}

function createServer() {
  return http.createServer(async (req, res) => {
    const url = new URL(req.url || '/', `http://${req.headers.host || 'localhost'}`);

    if (req.method !== 'GET' && req.method !== 'HEAD') {
      sendText(res, 405, 'Method not allowed');
      return;
    }

    if (url.pathname === '/') {
      await serveRoot(res);
      return;
    }

    if (url.pathname.startsWith('/static/')) {
      await serveStatic(url.pathname, res);
      return;
    }

    sendText(res, 404, 'Not found');
  });
}

if (require.main === module) {
  const server = createServer();
  server.listen(PORT, HOST, () => {
    process.stdout.write(`Server listening on http://${HOST}:${PORT}\n`);
  });
}

module.exports = {
  createServer
};
