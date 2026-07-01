export namespace main {
	
	export class UpdateCheckResult {
	    hasUpdate: boolean;
	    version: string;
	    url: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateCheckResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hasUpdate = source["hasUpdate"];
	        this.version = source["version"];
	        this.url = source["url"];
	        this.description = source["description"];
	    }
	}

}

