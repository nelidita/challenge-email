<script setup lang="ts">
import { RouterView } from 'vue-router'
import IconEmail from './components/icons/IconEmail.vue'
</script>

<template>
  <main class="h-full w-full overflow-hidden">
    <header>
      <nav class="flex gap-3 items-center font-bold text-xl text-slate-800 p-2">
        <IconEmail class="w-8 h-auto" />
        <p
          class="font-semibold text-base inline-flex animate-text-gradient bg-gradient-to-r from-[#ec8c3e] via-[#000000] to-[#faaf3f] bg-[200%_auto] bg-clip-text text-transparent"
        >
          Email Search
        </p>
      </nav>
    </header>
    <RouterView class="pt-12" />
  </main>
</template>
