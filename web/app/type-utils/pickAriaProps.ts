export type PickAriaProps<P extends object> = {
  [K in Extract<keyof P, `aria-${string}`>]: P[K]
}
