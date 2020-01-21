<template>
  <div class="menu-page">
    <div class="text-center">
      <div>
        {{ menuTitle }}
      </div>
      <div>
        <template v-for="(menu, menuIndex) in menuList">
          <router-link :to="{ name: menu.to }" :key="`menu-${menuIndex}`">
            <div>
              {{menu.title}}
            </div>
          </router-link>
        </template>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component';

@Component
export default class MenuPage extends Vue {
  menuTitle: string = 'MENU'
  menuList = [
    {
      title: 'ค้นหาผู้ป่วย',
      to: 'patient_search'
    },
    {
      title: 'เพิ่มข้อมูลผู้ป่วย',
      to: ''
    },
    {
      title: 'พิมพ์ใบรับรองแพทย์',
      to: ''
    },
    {
      title: 'ดูรายงานค่ารักษา',
      to: ''
    }
  ]
}
</script>