import type { RequestHandler } from '@sveltejs/kit';

const BASE = 'https://tehranbot.ir';

const routes = [
  '/',
];

export const GET: RequestHandler = async () => {
  const urls = routes
    .map((path) => `${BASE}${path}`)
    .map(
      (loc) => `  <url>
    <loc>${loc}</loc>
    <changefreq>weekly</changefreq>
    <priority>0.8</priority>
  </url>`
    )
    .join('\n');

  const body = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
${urls}
</urlset>`;

  return new Response(body, {
    headers: {
      'Content-Type': 'application/xml; charset=utf-8',
      'Cache-Control': 'public, max-age=3600'
    }
  });
};
