/*
 * Copyright (c) 2022 Institute of Software, Chinese Academy of Sciences (ISCAS)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index.js'
import installElementPlus from './plugins/element'
import 'default-passive-events'

import './style.scss'
import './assets/css/common.css'
import './assets/css/button.scss'
import './assets/css/font.css'

let app = createApp(App)
installElementPlus(app)
app.use(router).mount('#app')
