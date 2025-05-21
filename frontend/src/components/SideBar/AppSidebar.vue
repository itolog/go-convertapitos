<script setup lang="ts">
import { Home } from 'lucide-vue-next'
import NavMain from '@/components/SideBar/NavMain.vue'
import NavUser from '@/components/SideBar/NavUser.vue'

import {
  Sidebar,
  SidebarContent,
  SidebarHeader,
  SidebarRail,
  type SidebarProps,
} from '@/components/ui/sidebar'
import ThemeSwitcher from '@/components/ThemeSwitcher/ThemeSwitcher.vue'

const props = withDefaults(defineProps<SidebarProps>(), {
  collapsible: 'icon',
  variant: 'floating',
})

const data = {
  user: {
    name: '',
    email: '',
    avatar: '',
  },
  navMain: [
    {
      title: 'Home',
      url: '/',
      icon: Home,
    },
  ],
}
</script>

<template>
  <Sidebar v-bind="props">
    <SidebarHeader>
      <NavUser :user="data.user" />
    </SidebarHeader>
    <SidebarContent>
      <ThemeSwitcher />
      <NavMain :items="data.navMain" />
    </SidebarContent>
    <SidebarRail />
  </Sidebar>
</template>
