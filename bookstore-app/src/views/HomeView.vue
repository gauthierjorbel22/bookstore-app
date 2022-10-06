<template>
  <div class="home">
    <h3>Bookstore App</h3>
    <InsertForm />
  </div>
</template>

<script>
// @ is an alias to /src
import InsertForm from "@/components/InsertForm.vue";
export default {
  name: "HomeView",
  components: {
    InsertForm,
  },
};
</script>
