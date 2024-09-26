import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";
import { GameWebsocketProvider } from "@/context/game-socket";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "Nối từ",
  description: "Game nối từ",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <GameWebsocketProvider>
          {children}
        </GameWebsocketProvider>
      </body>
    </html>
  );
}
