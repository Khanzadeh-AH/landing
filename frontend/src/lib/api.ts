import { env as publicEnv } from '$env/dynamic/public';

export const API_BASE = (publicEnv.PUBLIC_API_BASE?.trim()) || 'http://localhost:8080/api';

export async function apiGet<T>(fetchFn: typeof fetch, path: string): Promise<T> {
  const res = await fetchFn(`${API_BASE}${path}`);
  if (!res.ok) throw new Error(`API ${path} failed: ${res.status}`);
  return res.json() as Promise<T>;
}
