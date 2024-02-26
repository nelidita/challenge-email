<script setup lang="ts">
const { isPage } = defineProps(['isPage'])
</script>

<template>
  <li>
    <button
      class="flex items-center justify-center transition duration-200 ease-in-out"
      :class="
        isPage
          ? 'font-extrabold px-1 h-full hover:text-orange-600'
          : 'h-7 w-7 rounded-full border border-gray-200 bg-transparent hover:bg-slate-200 active:bg-slate-300'
      "
    >
      <slot></slot>
    </button>
  </li>
</template>
