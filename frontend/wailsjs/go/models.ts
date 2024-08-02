export namespace apputil {
	
	export class MsgEmity {
	    success: boolean;
	    msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new MsgEmity(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}

}

