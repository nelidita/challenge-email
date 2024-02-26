<script setup lang="ts">
import { computed, ref } from 'vue'
import { randomColor } from '@/utils/randomColor'
import { formatDate } from '@/utils/formatDate'
import IconClose from './icons/IconClose.vue'

const props = defineProps(['modelValue', 'selectedEmail'])
const selectedEmail = ref(props.selectedEmail)
const emailBackground = computed(randomColor)
</script>

<template>
  <div class="font-sans bg-gray-100 flex items-center justify-center h-screen">
    <div class="fixed z-10 inset-0 flex items-center justify-center">
      <div class="absolute inset-0 bg-gray-500 opacity-75"></div>
      <div
        class="relative bg-white rounded-lg overflow-hidden shadow-xl max-w-screen-md w-full m-4"
      >
        <div
          class="prose max-w-screen-md p-4 overflow-y-auto flex flex-col gap-y-3"
          style="
            max-height: 70vh;
            background-color: #fff;
            border: 1px solid #e2e8f0;
            border-radius: 0.375rem;
            box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.1);
          "
        >
          <button
            type="button"
            class="fixed flex items-center self-end w-6 h-6 rounded-full bg-transparent transition-colors duration-300 active:bg-black/10 -mt-2 -mr-3"
            data-dismiss="modal"
            aria-label="Close"
            @click="$emit('closeModal')"
          >
            <IconClose />
          </button>
          <div class="flex flex-col gap-3 items-center w-full border-b-2 pt-5">
            <div class="flex self-start gap-2">
              <div
                class="flex items-center justify-center w-10 h-10 text-lg font-black shadow-sm shadow-current border border-slate-400"
                :style="{ backgroundColor: emailBackground, borderRadius: '100%' }"
              >
                {{ selectedEmail?.from.substring(0, 1) }}
              </div>
              <div class="w-fit">
                <p class="text-xs">From:</p>
                <p class="font-semibold truncate text-ellipsis overflow-hidden">
                  {{ selectedEmail?.from }}
                </p>
              </div>
            </div>
            <div class="flex flex-col text-base w-full gap-y-2">
              <div>
                <p class="text-xs">To:</p>
                <p class="font-semibold">{{ selectedEmail?.to }}</p>
              </div>
              <p class="text-light text-xs text-gray-400 pb-2">
                {{ formatDate(selectedEmail?.date) }}
              </p>
            </div>
          </div>
          <section>
            <h1 class="font-bold text-base">{{ selectedEmail?.subject }}</h1>
            <article class="mt-2 text-gray-500 leading-7 tracking-wider">
              <p>
                {{ selectedEmail?.content }}
              </p>
            </article>
          </section>
        </div>
      </div>
    </div>
  </div>
</template>
