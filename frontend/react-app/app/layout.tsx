import "./styles.css";

export const metadata = {
  title: "My NBA App",
  description: "My Cool NBA App",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
