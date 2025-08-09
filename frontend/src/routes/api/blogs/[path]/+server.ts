import type { RequestHandler } from './$types';
import { env } from '$env/dynamic/private';

export const GET: RequestHandler = async ({ fetch, params }) => {
  const backendBase = (env.BACKEND_API_BASE ?? 'http://localhost:8080/api').trim();
  const apiKey = (env.BACKEND_API_KEY ?? '').trim();

  const target = `${backendBase}/blogs/${encodeURIComponent(params.path)}`;

  const res = await fetch(target, {
    headers: apiKey ? { 'X-API-Key': apiKey } : undefined
  });

  const body = await res.text();
  return new Response(body, {
    status: res.status,
    headers: {
      'Content-Type': res.headers.get('Content-Type') || 'application/json'
    }
  });
};
