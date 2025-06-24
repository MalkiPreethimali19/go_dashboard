<template>
  <AppHeader />
  <DashboardView />
</template>

<script>
import DashboardView from './views/DashboardView.vue';
import AppHeader from './components/AppHeader.vue';
export default {
  components: {
    AppHeader, 
    DashboardView 
  }
};
</script>
