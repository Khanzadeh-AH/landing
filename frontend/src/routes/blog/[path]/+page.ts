import type { PageLoad } from './$types';
import { apiGet } from '$lib/api';

export type Blog = {
  id?: number;
  category: string;
  text: string;
  path: string;
};

export type BlogWithSimilar = { blog: Blog; similar: Blog[] };

export const load: PageLoad = async ({ fetch, params }) => {
  const res = await apiGet<BlogWithSimilar>(fetch, `/blogs/${encodeURIComponent(params.path)}`);
  return { blog: res.blog, similar: res.similar };
};
