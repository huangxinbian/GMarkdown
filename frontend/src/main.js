import Vue from "vue"
import App from "./App.vue"

import "./assects/iconfont/iconfont.css"

//快捷键
import {KeyMixin} from "./common/KeyMixin.js"
Vue.mixin(KeyMixin)

Vue.config.productionTip = false

import vuepressTheme from "./mavonEditor.js";

//打印插件
import Print from "vue-print-nb"
Vue.use(Print);

//-- wailsjs配置 开始 --//
Vue.prototype.go = {};

if (window.runtime) {
	for (let key in window.runtime) {
		if (window.runtime[key] instanceof Function) {
			Vue.prototype.go[key] = window.runtime[key];
		}
	}
}

if (window.go && window.go.main &&  window.go.main.App) {
	for (let key in window.go.main.App) {
		if (window.go.main.App[key] instanceof Function) {
			Vue.prototype.go[key] = window.go.main.App[key];
		}
	}
}

if (window.go && window.go.jsbind && window.go.jsbind.JsBind) {
	for (let key in window.go.jsbind.JsBind) {
		if (window.go.jsbind.JsBind[key] instanceof Function) {
			Vue.prototype.go[key] = window.go.jsbind.JsBind[key];
		}
	}
}
//-- wailsjs配置 结束 --//

new Vue({
	render: h => h(App),
}).$mount("#app")
