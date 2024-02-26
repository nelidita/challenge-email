<script setup lang="ts">
import { ref } from 'vue'
import { formatDateShort } from '@/utils/formatDate'
import type { Source } from '@/models/emails'

const { dataEmails } = defineProps(['dataEmails'])
const emit = defineEmits(['selectEmail'])
const idSelected = ref('')

const onSelectEmail = (dataEmail: Source) => {
  idSelected.value = dataEmail.id
  emit('selectEmail', dataEmail)
}
</script>

<template>
  <section
    class="w-full rounded-3xl sm:rounded-none sm:w-4/12 flex flex-col h-full bg-orange-50 border border-orange-200"
  >
    <div class="overflow-hidden">
      <ul class="overflow-y-auto h-full">
        <li
          v-for="email in dataEmails.hits"
          :key="email._id"
          @click="onSelectEmail(email._source)"
          class="py-5 border-b px-3 transition cursor-pointer rounded-3xl sm:rounded-none"
          :class="idSelected === email._source.id ? 'bg-orange-300' : 'hover:bg-slate-200'"
        >
          <div class="flex justify-between gap-x-1 items-center">
            <h3 class="text-sm font-semibold truncate text-ellipsis">{{ email._source.from }}</h3>
            <p class="text-xs text-end text-gray-400">{{ formatDateShort(email._source.date) }}</p>
          </div>
          <p class="text-xs text-gray-400 truncate text-ellipsis overflow-hidden">
            {{ email._source.subject }}
          </p>
        </li>
      </ul>
    </div>
  </section>
</template>
