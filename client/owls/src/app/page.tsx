import Link from 'next/link'
import styles from './page.module.css'



export default function Home() {
  return (
    <div>
    <main className={styles.main}>
      <Link href="/chat" className={styles.chat}>
        chat
      </Link>
    </main>
    </div>
  )
}
