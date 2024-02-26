<script setup lang="ts">
import { ref, onBeforeMount } from 'vue'
import { useEmailsStore } from '@/stores/emails'
import FilterComponent from '@/components/FilterComponent.vue'
import ListEmailsComponent from '@/components/ListEmailsComponent.vue'
import DetailEmail from '@/components/DetailEmailComponent.vue'
import LoadingComponent from '@/components/LoadingComponent.vue'
import PaginationComponent from '@/components/PaginationComponent.vue'
import ModalComponent from '@/components/ModalComponent.vue'
import type { Hits, Source } from '@/models/emails'
import type { IFilters } from '@/models/filters'
import { initialValues } from '@/constants/initialValues'

const emailsStore = useEmailsStore()
const data = ref<Hits>()
const filters = ref<IFilters>(initialValues)
const isLoading = ref<boolean>(false)
const selectedEmail = ref<Source>()
const currentPage = ref<number>(1)
const totalPages = ref<number>(0)
const showModal = ref(false)
const widthWindow = window.matchMedia('(max-width: 640px)').matches

onBeforeMount(async () => {
  isLoading.value = true
  await emailsStore.fetchEmails(initialValues)
  data.value = emailsStore.data
  totalPages.value = Math.ceil(emailsStore.data.total.value / 20)
  isLoading.value = false
  filters.value.orderBy = 'date'
  filters.value.order = 'asc'
})

const onSelectEmail = (email: Source) => {
  selectedEmail.value = email
  if (widthWindow) showModal.value = true
}

const handleSearch = async (filters: IFilters, page: number) => {
  selectedEmail.value = undefined
  currentPage.value = page
  const order = filters.order === 'asc' ? '' : '-'
  const sortFields = order + filters.orderBy

  const payload: IFilters = {
    wordSearch: filters.wordSearch,
    sortFields: [sortFields],
    page: page
  }

  await emailsStore.fetchEmails(payload)
  data.value = emailsStore.data
  totalPages.value = Math.ceil(emailsStore.data.total.value / 20)
}

const handleUpdatePage = async (page: number) => handleSearch(filters.value, page)

const handleCloseModal = () => {
  showModal.value = false
  selectedEmail.value = undefined
}
</script>

<template>
  <div class="w-screen h-screen flex">
    <div
      v-if="isLoading"
      class="w-screen h-screen flex flex-col gap-y-5 items-center justify-center"
    >
      <LoadingComponent />
    </div>
    <div v-else class="flex flex-col w-full h-full px-4 py-2">
      <PaginationComponent
        v-model="currentPage"
        :totalPages="totalPages"
        @updatePage="handleUpdatePage"
      />
      <div class="flex flex-col sm:flex-row h-[92%] drop-shadow-md gap-y-2 sm:gap-0">
        <FilterComponent v-model="filters" @search="handleSearch" />
        <ListEmailsComponent :dataEmails="data" @selectEmail="onSelectEmail" />
        <DetailEmail v-model="selectedEmail" />
        <ModalComponent
          v-if="showModal"
          v-model="showModal"
          :selectedEmail="selectedEmail"
          @closeModal="handleCloseModal"
        />
      </div>
    </div>
  </div>
</template>
