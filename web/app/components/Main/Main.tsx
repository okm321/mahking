import type { PropsWithChildren } from "react";
import styles from './Main.module.scss';

export function Main({
  children
}: PropsWithChildren) {
  return (
    <div className={styles.main_wrapper}>
      <main className={styles.main}>
        {children}
      </main>
    </div>
  )
}
