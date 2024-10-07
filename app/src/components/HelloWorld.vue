<script setup lang="ts">
defineProps<{
  msg: string;
}>();
</script>

<template>
  <div class="greetings">
    <h1 class="green">{{ msg }}</h1>
    <!-- <h3>
      AM I WORKING NOW?
    </h3> -->
  </div>
</template>
