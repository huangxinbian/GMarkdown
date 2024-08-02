export const KeyMixin = {
	data() {
		return {
			respond: false
		};
	},
	methods: {
		// 键盘抬起
		keyup(event) {
			this.respond = false;
		},
		
		// 键盘按下
		keyDown(event, shortcutArray) {
			if (this.respond) {
				return;
			}
			
			if ((!event.altKey) && (!event.ctrlKey)) {
				return;
			}
			
			let msg = event.altKey ? "alt " : "";
			msg = msg + (event.ctrlKey ? "ctrl " : "");
			msg = msg + event.key.toUpperCase();//event.key = "t"//"T"
			
			let keyInfo = {
				key: event.key.toUpperCase(),
				alt: event.altKey,
				ctrl: event.ctrlKey,
				msg: msg
			}
			
			let fun = this.valid(shortcutArray, keyInfo);
			if (!fun) {
				return;
			}
			
			event.preventDefault();
			this.respond = true;
			fun(event);
		},
		
		valid(shortcutArray, keyInfo) {
			for (let i = 0; i < shortcutArray.length; i++) {
				let item = shortcutArray[i];
				if (item.key.toUpperCase() != keyInfo.key) {
					continue;
				}
				
				if (keyInfo.alt && !item.alt) {
					continue;
				}
				
				if (keyInfo.ctrl && !item.ctrl) {
					continue;
				}
				
				if (item.fun) {
					return item.fun;
				}
			}
			
			return null;
		}
	},
	//--加载后执行--//
	mounted() {
	}
};