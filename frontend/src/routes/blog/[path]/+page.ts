import type { PageLoad } from './$types';
import { apiGet } from '$lib/api';

export type Blog = {
  id?: number;
  category: string;
  text: string;
  path: string;
};

export const load: PageLoad = async ({ fetch, params }) => {
  const blog = await apiGet<Blog>(fetch, `/blogs/${encodeURIComponent(params.path)}`);
  return { blog };
};
