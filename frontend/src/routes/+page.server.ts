import { fail, type Actions } from '@sveltejs/kit';
import { z } from 'zod';
import nodemailer from 'nodemailer';

const schema = z.object({
  name: z.string({ required_error: 'نام الزامی است' }).min(3, 'نام حداقل ۳ کاراکتر باشد'),
  email: z.string({ required_error: 'ایمیل الزامی است' }).email('ایمیل معتبر نیست'),
  phone: z
    .string()
    .optional()
    .refine((v) => !v || /^\+?[0-9\s-]{7,15}$/.test(v), { message: 'شماره تماس معتبر نیست' }),
  message: z.string({ required_error: 'پیام الزامی است' }).min(10, 'لطفاً پیام خود را کامل‌تر شرح دهید')
});

export const actions: Actions = {
  default: async ({ request }) => {
    const formData = await request.formData();
    const raw = {
      name: formData.get('name')?.toString() ?? '',
      email: formData.get('email')?.toString() ?? '',
      phone: formData.get('phone')?.toString() ?? '',
      message: formData.get('message')?.toString() ?? ''
    };

    const parsed = schema.safeParse(raw);
    if (!parsed.success) {
      const errors: Record<string, string> = {};
      for (const issue of parsed.error.issues) {
        const field = String(issue.path[0]);
        if (!errors[field]) errors[field] = issue.message;
      }
      return fail(400, { errors, values: raw });
    }

    try {
      const { SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASS, SMTP_FROM, CONTACT_TO } = process.env;
      if (SMTP_HOST && SMTP_USER && SMTP_PASS && CONTACT_TO) {
        const transporter = nodemailer.createTransport({
          host: SMTP_HOST,
          port: Number(SMTP_PORT || 587),
          secure: Number(SMTP_PORT || 587) === 465,
          auth: { user: SMTP_USER, pass: SMTP_PASS }
        });

        await transporter.sendMail({
          from: `TehranBot Website <${SMTP_FROM || SMTP_USER}>`,
          to: CONTACT_TO,
          subject: 'پیام جدید از فرم تماس تهران‌بات',
          text: `نام: ${raw.name}\nایمیل: ${raw.email}\nشماره: ${raw.phone}\n\n${raw.message}`,
          html: `<p><strong>نام:</strong> ${raw.name}</p><p><strong>ایمیل:</strong> ${raw.email}</p><p><strong>شماره:</strong> ${raw.phone}</p><p>${raw.message.replace(/\n/g, '<br>')}</p>`
        });
      } else {
        console.log('Contact form submission', raw);
      }
    } catch (err) {
      console.error('Failed to send contact email', err);
      // Do not block success on email errors
    }

    return { success: true };
  }
};
