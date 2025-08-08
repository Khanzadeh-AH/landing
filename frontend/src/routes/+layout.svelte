<script lang="ts">
    import '../app.css';
    import favicon from '$lib/assets/favicon.svg';

    let { children } = $props();

    const orgLd = {
        '@context': 'https://schema.org',
        '@type': 'Organization',
        name: 'TehranBot',
        alternateName: 'تهران‌بات',
        url: 'https://tehranbot.ir/',
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
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
    <title>تهران‌بات | توسعه نرم‌افزار و خدمات هوش مصنوعی</title>
    <meta name="description" content="طراحی و توسعه نرم‌افزارهای تحت وب و موبایل، ساخت ربات‌های هوشمند، RAG، DevOps و اتوماسیون فرآیندها. مشاوره رایگان برای شروع پروژه شما." />
    <meta name="theme-color" content="#06b6d4" />
    <meta property="og:type" content="website" />
    <meta property="og:locale" content="fa_IR" />
    <meta property="og:title" content="تهران‌بات | توسعه نرم‌افزار و خدمات هوش مصنوعی" />
    <meta property="og:description" content="خدمات تخصصی توسعه نرم‌افزار و هوش مصنوعی برای کسب‌وکارها" />
    <meta property="og:site_name" content="TehranBot" />
    <script type="application/ld+json">{JSON.stringify(orgLd)}</script>
</svelte:head>

{@render children?.()}
