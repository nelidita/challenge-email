import { defineStore } from 'pinia'
import type { IFilters } from '../models/filters'
import type { IEmails } from '../models/emails'
import type { Hits } from '../models/emails'

export const useEmailsStore = defineStore('emails', {
  state: (): IEmails => ({ data: {} as Hits }),
  actions: {
    async fetchEmails(payload: IFilters) {
      const response = await fetch(`${import.meta.env.VITE_API_URL}emails/search`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          word_search: payload.wordSearch,
          sort_fields: payload.sortFields,
          page: payload.page
        })
      })
      const { hits } = await response.json()
      this.data = hits
    }
  }
})
