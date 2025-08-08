<section id="about" class="section">
  <div class="container-rtl">
    <div class="text-center">
      <h2 class="text-3xl md:text-4xl">چرا تهران‌بات؟</h2>
      <p class="lead mt-2">ترکیبی از تجربه، تخصص و تعهد</p>
    </div>

    <div class="mt-8 grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
      {#each [
        'تجربه در پروژه‌های متنوع',
        'استفاده از جدیدترین فناوری‌ها',
        'پشتیبانی ۲۴/۷',
        'تحویل به‌موقع پروژه‌ها',
        'قیمت‌گذاری منصفانه و شفاف',
        'تخصص فنی بالا'
      ] as item}
        <div class="card p-6 flex items-center gap-3">
          <svg class="w-6 h-6 text-primary-600 shrink-0" viewBox="0 0 20 20" fill="currentColor" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/></svg>
          <span>{item}</span>
        </div>
      {/each}
    </div>
  </div>
</section>
