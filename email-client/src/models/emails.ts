export interface IEmails {
  data: Hits
}

export interface IEmail {
  id: string
  from: string
  to: string
  content: string
  subject: string
  date: string
  filepath: string
}

export interface Hits {
  hits: Hit[]
  total: Total
  max_score: number
}

export interface Hit {
  _id: string
  '@timestamp': Date
  _score: number
  _source: Source
}

export interface Source {
  '@timestamp': Date
  content: string
  date: string
  filepath: string
  from: Date
  id: string
  subject: string
  to: string
}

export interface Total {
  value: number
}
