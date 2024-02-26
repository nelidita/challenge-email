export function randomColor(): string {
  const r = Math.floor(Math.random() * 256)
  const g = Math.floor(Math.random() * 256)
  const b = Math.floor(Math.random() * 256)
  const a = Math.random()
  return `rgba(${r}, ${g}, ${b}, ${a})`
}
