<template>
  <div>
    <header
      class="bg-cyan-500 dark:bg-gray-800 grid grid-cols-2 md:py-[.2em] w-full z-30 will-change-transform"
    >
      <div class="px-[2em] py-[1em] block">
        <nuxt-link to="/" class="uppercase font-bold text-white"
          >Bottlehub</nuxt-link
        >
      </div>

      <nav class="justify-self-end">
        <div class="md:hidden flex items-center">
          <button
            class="gradient-button mobile-menu-button"
            @click="showMenu = !showMenu"
          >
            <div class="outline-none p-1 rounded bg-cyan-600 dark:bg-gray-900">
              <svg
                class="w-7 h-7 text-gray-500"
                x-show="!shoMenu"
                fill="none"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path d="M4 6h16M4 12h16M4 18h16"></path>
              </svg>
            </div>
          </button>
        </div>
        <ul class="hidden md:flex">
          <li class="gradient-li">
            <nuxt-link to="/" class="partials">Home</nuxt-link>
          </li>
          <li class="gradient-li">
            <nuxt-link to="#about" class="partials">About</nuxt-link>
          </li>
          <li class="gradient-li">
            <nuxt-link to="#support" class="partials">Support</nuxt-link>
          </li>
          <!-- <li
            class="gradient-li"
          >
            <nuxt-link
              to="#contact"
              class="partials"
              >Contact</nuxt-link
            >
          </li> -->
        </ul>
      </nav>
    </header>
    <div
      class="animate-border bg-[length:200%_200%] hover:bg-[length:400%_400%] grid grid-cols-2 md:py-[.15em] bg-gradient-to-r from-red-600 via-purple-600 to-blue-600 w-full z-30 will-change-transform"
    ></div>
    <div class="mobile-menu w-full md:hidden right-0 bg-slate-600" :class="showMenu ? 'absolute' : 'hidden'">
      <ul class="text-center">
        <li>
          <nuxt-link
            to="/"
            class="block text-sm py-4 hover:animate-border hover:bg-[length:200%_200%] hover:bg-gradient-to-r from-red-600 via-purple-600 to-blue-600"
            >Home</nuxt-link
          >
        </li>
        <li>
          <nuxt-link
            to="#about"
            class="block text-sm py-4 hover:animate-border hover:bg-[length:200%_200%] hover:bg-gradient-to-r from-red-600 via-purple-600 to-blue-600"
            >About</nuxt-link
          >
        </li>
        <li>
          <nuxt-link
            to="#support"
            class="block text-sm py-4 hover:animate-border hover:bg-[length:200%_200%] hover:bg-gradient-to-r from-red-600 via-purple-600 to-blue-600"
            >Support</nuxt-link
          >
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showMenu: false,
    };
  },
};
</script>

<style>
.gradient-li {
  @apply inline hover:animate-border bg-[length:200%_200%] hover:bg-[length:400%_400%] m-1 px-[.025em] py-[.025em] rounded-full max-w-sm bg-gradient-to-r from-red-600 via-purple-600 to-blue-600;
}
.gradient-button {
  @apply inline hover:animate-border bg-[length:200%_200%] hover:bg-[length:400%_400%] m-1 p-[.2em] rounded max-w-sm bg-gradient-to-r from-red-600 via-purple-600 to-blue-600;
}
.gradient-button:focus-within {
  @apply animate-border bg-[length:400%_400%];
}
.partials {
  @apply px-[2em] py-[1em] inline-block uppercase font-bold text-white text-sm bg-cyan-600 dark:bg-gray-900 mx-[.3em] my-1 rounded-[2rem];
}
</style>
