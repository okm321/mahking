import { Stack } from '~/components/shared/Stack'
import styles from './HomeView.module.scss'
import { Button } from '~/components/shared/Button'
import { SquarePlus, Zap, ChartColumn, Share2 } from 'lucide-react'
import { Link } from '@tanstack/react-router'

const features = [
  { icon: Zap, title: 'かんたん点数入力', description: '素点を入れるだけで自動計算' },
  { icon: ChartColumn, title: 'グラフで見える成績', description: 'スコア・着順の推移がひと目で' },
  { icon: Share2, title: 'URLでみんなと共有', description: '登録不要、リンクを送るだけ' },
] as const

export function HomeView() {
  return (
    <Stack spacing={4} className={styles.home}>
      <h2 className={styles.home_title}>麻雀の記録を<br />手軽にシンプルに</h2>
      <p className={styles.home_description}>ログイン不要。URLを共有するだけで、みんなの成績を記録できます。さまざまなルールにも柔軟に対応しています。</p>
      <Button size="l" bold startIcon={<SquarePlus />} as={Link} to="/new" className={styles['start-button']}>今すぐ始める</Button>
      <ul className={styles.feature_list}>
        {features.map(({ icon: Icon, title, description }) => (
          <li key={title} className={styles.feature_item}>
            <span className={styles.feature_icon} aria-hidden="true"><Icon size={20} /></span>
            <div>
              <p className={styles.feature_title}>{title}</p>
              <p className={styles.feature_description}>{description}</p>
            </div>
          </li>
        ))}
      </ul>
    </Stack>
  )
}
