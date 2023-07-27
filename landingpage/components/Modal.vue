<template>
  <button @click="open = true">Open Modal</button>
  <Teleport to="body">
    <div v-if="open" class="modal-bg w-[100%] h-[100vh] flex justify-center items-center">
      <div class="modal bg-slate-600 rounded-md p-4 w-[300px]">
        <p>Hello from the modal!</p>
        <button @click="open = false">Close</button>
      </div>
    </div>
  </Teleport>
</template>

<script>
export default {
  data() {
    return {
      open: false,
    };
  },
};
</script>
