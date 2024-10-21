export namespace models {
	
	export class Config {
	    Status: number;
	    Port: number;
	    AutoProxy: boolean;
	    SaveLogFile: boolean;
	    Filter: boolean;
	    FilterHost: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Status = source["Status"];
	        this.Port = source["Port"];
	        this.AutoProxy = source["AutoProxy"];
	        this.SaveLogFile = source["SaveLogFile"];
	        this.Filter = source["Filter"];
	        this.FilterHost = source["FilterHost"];
	    }
	}
	export class SitemapConfig {
	    Dsturl: string;
	    Savepath: string;
	    Filename: string;
	    Concurrency: number;
	    Crawltimeout: number;
	    Requesttimeout: number;
	
	    static createFrom(source: any = {}) {
	        return new SitemapConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Dsturl = source["Dsturl"];
	        this.Savepath = source["Savepath"];
	        this.Filename = source["Filename"];
	        this.Concurrency = source["Concurrency"];
	        this.Crawltimeout = source["Crawltimeout"];
	        this.Requesttimeout = source["Requesttimeout"];
	    }
	}

}

