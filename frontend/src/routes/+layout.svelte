<script lang="ts">
    import '../app.css';
    import favicon from '$lib/assets/favicon.svg';

    let { children } = $props();

    const orgLd = {
        '@context': 'https://schema.org',
        '@type': 'Organization',
        name: 'TehranBot',
        alternateName: 'تهران‌بات',
        url: 'https://tehranbot.me/',
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
        name: 'TehranBot',
        image: 'https://tehranbot.me/og-image.svg',
        url: 'https://tehranbot.me/',
        telephone: '+98 919 073 7241',
        address: {
            '@type': 'PostalAddress',
            addressLocality: 'Tehran',
            addressCountry: 'IR'
        }
    } as const;
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
    <link rel="canonical" href="https://tehranbot.me/" />
    <link rel="alternate" hreflang="fa-IR" href="https://tehranbot.me/" />
    <title>تهران‌بات | توسعه نرم‌افزار و خدمات هوش مصنوعی</title>
    <meta name="description" content="طراحی و توسعه نرم‌افزارهای تحت وب و موبایل، ساخت ربات‌های هوشمند، RAG، DevOps و اتوماسیون فرآیندها. مشاوره رایگان برای شروع پروژه شما." />
    <meta name="theme-color" content="#06b6d4" />
    <meta name="robots" content="index,follow" />
    <meta property="og:type" content="website" />
    <meta property="og:locale" content="fa_IR" />
    <meta property="og:url" content="https://tehranbot.me/" />
    <meta property="og:title" content="تهران‌بات | توسعه نرم‌افزار و خدمات هوش مصنوعی" />
    <meta property="og:description" content="خدمات تخصصی توسعه نرم‌افزار و هوش مصنوعی برای کسب‌وکارها" />
    <meta property="og:site_name" content="TehranBot" />
    <meta property="og:image" content="https://tehranbot.me/og-image.svg" />
    <meta property="og:image:width" content="1200" />
    <meta property="og:image:height" content="630" />
    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:title" content="تهران‌بات | توسعه نرم‌افزار و خدمات هوش مصنوعی" />
    <meta name="twitter:description" content="خدمات تخصصی توسعه نرم‌افزار و هوش مصنوعی برای کسب‌وکارها" />
    <meta name="twitter:image" content="https://tehranbot.me/og-image.svg" />
    <meta name="twitter:image:alt" content="تهران‌بات | توسعه نرم‌افزار و خدمات هوش مصنوعی" />
    <script type="application/ld+json">{JSON.stringify(orgLd)}</script>
    <script type="application/ld+json">{JSON.stringify(localLd)}</script>
</svelte:head>

<a href="#main-content" class="sr-only focus:not-sr-only fixed top-2 right-2 z-50 rounded-lg bg-primary-600 px-3 py-2 text-white">پرش به محتوا</a>

{@render children?.()}
