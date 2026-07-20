import { Link } from '@tanstack/react-router'
import styles from './Header.module.scss'

export function Header() {
  return (
    <header className={styles.header}>
      <h1 className={styles.header__title}>
        <Link to="/">Mahking</Link>
      </h1>
    </header>
  )
}
