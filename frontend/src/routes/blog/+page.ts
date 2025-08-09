import type { PageLoad } from './$types';
import { apiGet } from '$lib/api';

export type Blog = {
  id?: number;
  category: string;
  text: string;
  path: string;
  // ent typically adds created_at/updated_at if defined; we didn't, so keep minimal
};

export const load: PageLoad = async ({ fetch, url }) => {
  const category = url.searchParams.get('category')?.trim() || '';
  const qs = category ? `?category=${encodeURIComponent(category)}` : '';
  const blogs = await apiGet<Blog[]>(fetch, `/blogs${qs}`);
  return { blogs, category };
};
