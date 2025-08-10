<script lang="ts">
    import '../app.css';
    import favicon from '$lib/assets/favicon.svg';
    import Header from '$lib/components/Header.svelte';
    import { env as publicEnv } from '$env/dynamic/public';

    let { children } = $props();

    const BASE = (publicEnv.PUBLIC_BASE_URL?.trim()) || 'https://tehranbot.me';
    const SITE = {
        name: 'TehranBot',
        altName: 'تهران‌بات',
        url: `${BASE}/`,
        image: `${BASE}/og-image.svg`,
        title: 'تهران‌بات | توسعه نرم‌افزار و خدمات هوش مصنوعی',
        description: 'طراحی و توسعه نرم‌افزارهای تحت وب و موبایل، ساخت ربات‌های هوشمند، RAG، DevOps و اتوماسیون فرآیندها. مشاوره رایگان برای شروع پروژه شما.'
    } as const;

    const orgLd = {
        '@context': 'https://schema.org',
        '@type': 'Organization',
        name: SITE.name,
        alternateName: SITE.altName,
        url: SITE.url,
        email: 'mailto:khanzadeh78ah@gmail.com',
        sameAs: [
            'https://t.me/Khanzadeh_AH'
        ],
        contactPoint: [
            {
                '@type': 'ContactPoint',
                telephone: '+98 919 073 7241',
                contactType: 'customer service',
                areaServed: 'IR',
                availableLanguage: ['fa']
            }
        ]
    } as const;

    const localLd = {
        '@context': 'https://schema.org',
        '@type': 'LocalBusiness',
        name: SITE.name,
        image: SITE.image,
        url: SITE.url,
        telephone: '+98 919 073 7241',
        address: {
            '@type': 'PostalAddress',
            addressLocality: 'Tehran',
            addressCountry: 'IR'
        }
    } as const;

    const orgLdStr = JSON.stringify(orgLd);
    const localLdStr = JSON.stringify(localLd);
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
    <link rel="canonical" href={SITE.url} />
    <link rel="alternate" hreflang="fa-IR" href={SITE.url} />
    <title>{SITE.title}</title>
    <meta name="description" content={SITE.description} />
    <meta name="theme-color" content="#06b6d4" />
    <meta name="robots" content="index,follow" />
    <meta property="og:type" content="website" />
    <meta property="og:locale" content="fa_IR" />
    <meta property="og:url" content={SITE.url} />
    <meta property="og:title" content={SITE.title} />
    <meta property="og:description" content={SITE.description} />
    <meta property="og:site_name" content={SITE.name} />
    <meta property="og:image" content={SITE.image} />
    <meta property="og:image:width" content="1200" />
    <meta property="og:image:height" content="630" />
    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:title" content={SITE.title} />
    <meta name="twitter:description" content={SITE.description} />
    <meta name="twitter:image" content={SITE.image} />
    <meta name="twitter:image:alt" content={SITE.title} />
    <script type="application/ld+json">{JSON.stringify(orgLd)}</script>
    <script type="application/ld+json">{JSON.stringify(localLd)}</script>
</svelte:head>

<a href="#main-content" class="sr-only focus:not-sr-only fixed top-2 right-2 z-50 rounded-lg bg-primary-600 px-3 py-2 text-white">پرش به محتوا</a>

<!-- Site-wide decorative background -->
<div aria-hidden="true" class="pointer-events-none fixed inset-0 -z-10">
    <div class="absolute -top-24 -right-20 h-[40vh] w-[40vh] rounded-full bg-primary-500/10 blur-3xl"></div>
    <div class="absolute -bottom-24 -left-20 h-[50vh] w-[50vh] rounded-full bg-cyan-400/10 blur-[80px]"></div>
</div>

<Header />
{@render children?.()}
