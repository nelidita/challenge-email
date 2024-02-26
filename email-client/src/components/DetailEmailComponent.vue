<script setup lang="ts">
import { watch, ref, computed } from 'vue'
import IconEmailEmpty from '@/components/icons/IconEmailEmpty.vue'
import { randomColor } from '@/utils/randomColor'
import { formatDate } from '@/utils/formatDate'
import type { Source } from '@/models/emails'

const props = defineProps(['modelValue'])
const selectedEmail = ref<Source>()

const emailBackground = computed(() => {
  if (!selectedEmail.value) return ''
  return randomColor()
})

watch(
  () => props.modelValue,
  () => (selectedEmail.value = props.modelValue)
)
</script>

<template>
  <section
    class="hidden sm:flex w-6/12 h-full flex-col overflow-hidden bg-orange-50 rounded-r-3xl border border-orange-200 border-l-0"
  >
    <div v-if="selectedEmail" class="overflow-y-auto h-full px-4">
      <div class="flex gap-x-3 items-center pb-2 pt-4 w-full border-b-2 mb-4">
        <div class="w-1/12 flex items-center justify-center">
          <div
            class="flex items-center justify-center w-10 h-10 text-lg font-black shadow-sm shadow-current border border-slate-400"
            :style="{ backgroundColor: emailBackground, borderRadius: '100%' }"
          >
            {{ selectedEmail?.from.substring(0, 1) }}
          </div>
        </div>
        <div class="flex flex-col text-base w-11/12 gap-y-2">
          <div>
            <p class="text-xs">From:</p>
            <p class="font-semibold">{{ selectedEmail?.from }}</p>
          </div>
          <div>
            <p class="text-xs">To:</p>
            <p class="font-semibold">{{ selectedEmail?.to }}</p>
          </div>
          <p class="text-light text-xs text-gray-400">{{ formatDate(selectedEmail?.date) }}</p>
        </div>
      </div>
      <section>
        <h1 class="font-bold text-base">{{ selectedEmail?.subject }}</h1>
        <article class="mt-4 text-gray-500 leading-7 tracking-wider">
          <p>
            {{ selectedEmail?.content }}
          </p>
        </article>
      </section>
    </div>
    <div v-else class="flex flex-col items-center justify-center gap-y-3 h-full">
      <IconEmailEmpty class="w-48" />
      <p class="text-base text-gray-500">Select an email to see its details</p>
    </div>
  </section>
</template>
