import { format } from 'date-fns'

export function formatDate(dateString?: string) {
  if (dateString) {
    const date = new Date(dateString.replace(/\s-\d{4}.*$/, ''))
    const formattedDate = format(date, 'ccc, dd MMM y h:mm a')
    return formattedDate
  }
}

export function formatDateShort(dateEmail: string) {
  const date = new Date(dateEmail)
  const day = date.getDate()
  const month = date.toLocaleString('default', { month: 'short' })
  const year = date.getFullYear()
  return `${day} ${month}, ${year}`
}
