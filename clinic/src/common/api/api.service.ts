import axios, { AxiosResponse } from 'axios';
import Vue from "vue";
import VueAxios from "vue-axios";
import { API_URL } from './config';

const ApiService = {
  init() {
    Vue.use(VueAxios,axios);
    Vue.axios.defaults.baseURL = API_URL;
  },
  post<R>(apiPath: string, request: any): Promise<AxiosResponse<R>> {
    return Vue.axios.post<R>(apiPath, request);
  }

}

export default ApiService
