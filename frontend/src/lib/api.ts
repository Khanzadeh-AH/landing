export const API_BASE = '/api';

export async function apiGet<T>(fetchFn: typeof fetch, path: string): Promise<T> {
  const res = await fetchFn(`${API_BASE}${path}`);
  if (!res.ok) throw new Error(`API ${path} failed: ${res.status}`);
  return res.json() as Promise<T>;
}
