<script setup lang="ts">
import { ref, watch } from 'vue'
import IconArrowLeft from './icons/IconArrowLeft.vue'
import IconArrowRight from './icons/IconArrowRight.vue'
import IconArrowStart from './icons/IconArrowStart.vue'
import IconArrowEnd from './icons/IconArrowEnd.vue'
import ButtonPagComponent from './ButtonPagComponent.vue'

const props = defineProps(['modelValue', 'totalPages'])
const emit = defineEmits(['updatePage'])
const currentPage = ref(props.modelValue)

const handleUpdatePage = (page: number) => {
  if (page >= 1 && page <= props.totalPages) {
    currentPage.value = page
    emit('updatePage', page)
  }
}

watch(
  () => props.modelValue,
  (newValue) => (currentPage.value = newValue)
)
</script>

<template>
  <div class="grid h-[8%] w-full place-items-end overflow-x-scroll lg:overflow-visible pb-1">
    <nav>
      <ul class="flex gap-x-1">
        <ButtonPagComponent v-if="currentPage > 1" @click="handleUpdatePage(1)" :isPage="false">
          <IconArrowStart class="w-3" />
        </ButtonPagComponent>
        <ButtonPagComponent
          v-if="currentPage > 1"
          @click="handleUpdatePage(currentPage - 1)"
          :isPage="false"
        >
          <IconArrowLeft class="w-3" />
        </ButtonPagComponent>
        <ButtonPagComponent v-if="currentPage > 1" :isPage="true">
          {{ currentPage - 1 }}
        </ButtonPagComponent>
        <ButtonPagComponent :isPage="true" class="text-lg border-b-2 border-b-orange-400">
          {{ currentPage }}
        </ButtonPagComponent>
        <ButtonPagComponent v-if="currentPage < totalPages" :isPage="true">
          {{ currentPage + 1 }}
        </ButtonPagComponent>
        <ButtonPagComponent
          v-if="currentPage < totalPages"
          @click="handleUpdatePage(currentPage + 1)"
          :isPage="false"
        >
          <IconArrowRight class="w-3" />
        </ButtonPagComponent>
        <ButtonPagComponent
          v-if="currentPage < totalPages"
          @click="handleUpdatePage(totalPages)"
          :isPage="false"
        >
          <IconArrowEnd class="w-3" />
        </ButtonPagComponent>
      </ul>
    </nav>
  </div>
</template>
