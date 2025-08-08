export function reveal(node: HTMLElement, options: { once?: boolean } = { once: true }) {
  const prefersReduced = globalThis.matchMedia?.('(prefers-reduced-motion: reduce)').matches;
  if (prefersReduced) return {};

  node.style.transition = 'opacity 400ms ease, transform 400ms ease';
  node.style.willChange = 'opacity, transform';
  node.style.opacity = '0';
  node.style.transform = 'translateY(12px)';

  const observer = new IntersectionObserver(
    (entries) => {
      for (const entry of entries) {
        if (entry.isIntersecting) {
          node.style.opacity = '1';
          node.style.transform = 'translateY(0)';
          if (options.once !== false) observer.unobserve(node);
        }
      }
    },
    { threshold: 0.15 }
  );

  observer.observe(node);

  return {
    destroy() {
      observer.disconnect();
    }
  };
}
