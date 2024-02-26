<script setup lang="ts">
import { ref } from 'vue'
import IconSearch from './icons/IconSearch.vue'

const props = defineProps(['modelValue'])
const filters = ref(props.modelValue)
const initPage = 1
</script>

<template>
  <form
    @submit.prevent="$emit('search', filters, initPage)"
    class="flex flex-col gap-2 sm:gap-0 w-full sm:w-2/12 justify-between h-full bg-orange-50 rounded-3xl sm:rounded-r-none p-4 border border-orange-200 border-r-0"
  >
    <div class="flex flex-row sm:flex-col gap-2">
      <p class="hidden sm:flex font-bold text-base">Filters:</p>
      <div class="relative drop-shadow-md hover:drop-shadow-lg">
        <label for="sort" class="flex sm:hidden text-sm pb-1">Search:</label>
        <input
          type="text"
          class="w-full p-1 pl-7 text-sm rounded-lg border border-orange-400 bg-orange-200 focus:bg-white focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent hover:border-orange-600"
          placeholder="Search..."
          v-model="filters.wordSearch"
        />
        <IconSearch class="w-4 absolute left-2.5 top-8 sm:top-2" />
      </div>
      <div class="flex flex-col gap-1 drop-shadow-md hover:drop-shadow-lg">
        <label for="sort" class="text-sm">Order by:</label>
        <select
          id="sort"
          v-model="filters.orderBy"
          class="p-1 rounded-md text-sm border border-orange-400 bg-orange-200 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent hover:border-orange-600"
        >
          <option value="from">From</option>
          <option value="to">To</option>
          <option value="subject">Subject</option>
          <option value="date">Date</option>
          <option value="content">Content</option>
        </select>
      </div>
      <div class="flex flex-col gap-1 drop-shadow-md hover:drop-shadow-lg">
        <label for="order" class="text-sm">Order:</label>
        <select
          id="order"
          v-model="filters.order"
          class="p-1 rounded-md text-sm border border-orange-400 bg-orange-200 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent hover:border-orange-600"
        >
          <option value="asc" data-icon="asc">Asc</option>
          <option value="desc" data-icon="desc">Desc</option>
        </select>
      </div>
    </div>
    <button
      type="submit"
      class="inline-flex animate-background-shine items-center justify-center rounded-full border border-orange-400 bg-[linear-gradient(110deg,#fda619,45%,#fff38e,55%,#fda619)] bg-[length:200%_100%] p-1 text-sm font-semibold transition-colors active:outline-none active:ring-2 active:ring-orange-400 active:ring-offset-2 active:ring-offset-orange-50 shadow-md shadow-orange-400 hover:ring-2 hover:ring-orange-400 hover:outline-none hover:ring-offset-1"
    >
      Search
    </button>
  </form>
</template>
