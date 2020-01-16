import axios from 'axios';
import Vue from "vue";
import VueAxios from "vue-axios";
import { API_URL } from './config';

const ApiService = {
  init(){
    Vue.use(VueAxios,axios);
    Vue.axios.defaults.baseURL = API_URL;
  },
  post(apiPath: string, body: any) {
    return Vue.axios.post(apiPath, body);
  }

}

export default ApiService
