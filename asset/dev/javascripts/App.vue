<template>
    <div class="root">
        <div class="row">
            <div class="col-9">
                <div class="form-group">
                    <label for="jira">JIRA Task</label>
                    <input type="text" class="form-control form-control-sm" id="jira" aria-describedby="jiraHelp"
                           placeholder="and only the task number (including 'FU-')" v-model="jira"
                           v-on:change="jiraChange">
                    <small id="jiraHelp" class="form-text text-muted">Seperate by comma (,)</small>
                </div>
                <div class="form-group">
                    <label for="linkedPr">Linked P.R.</label>
                    <input type="text" class="form-control form-control-sm" id="linkedPr"
                           aria-describedby="linkedPrHelp" v-model="linkedPr" v-on:change="linkedPrChange">
                    <small id="linkedPrHelp" class="form-text text-muted">Seperate by comma (,)</small>
                </div>
                <div class="form-group">
                    <label for="relatedPr">Related P.R.</label>
                    <input type="text" class="form-control form-control-sm" id="relatedPr"
                           aria-describedby="relatedPrHelp" v-model="relatedPr" v-on:change="relatedPrPrChange">
                    <small id="relatedPrHelp" class="form-text text-muted">Seperate by comma (,)</small>
                </div>
            </div>
            <div class="col-3">
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="CMS" v-model="cms" v-on:change="csmChange">
                    <label class="form-check-label" for="CMS">CMS</label>
                </div>
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="DB" v-model="db" v-on:change="dbChange">
                    <label class="form-check-label" for="DB">DB</label>
                </div>
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="LiveFollowUp" v-model="liveFollowUp"
                           v-on:change="liveFollowUpChange">
                    <label class="form-check-label" for="LiveFollowUp">Live Follow-up</label>
                </div>
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="Parameters" v-model="parameters"
                           v-on:change="parametersChange">
                    <label class="form-check-label" for="Parameters">Parameters</label>
                </div>
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="Cloudinary" v-model="cloudinary"
                           v-on:change="cloudinaryChange">
                    <label class="form-check-label" for="Cloudinary">Requires cloudinary upload</label>
                </div>

                <span class="btn btn-secondary btn-sm mt-2" id="clear" v-on:click="clearField"><i
                        class="fas fa-broom"></i> Clear Fields</span>
            </div>
        </div>

        <div class="row mt-4">
            <div class="col-9">
                <h2>Output</h2>
            </div>
            <div class="col-3 text-right">
                <span :class="clipBoardStatus" v-on:click="clipboard">
                    <i class="icon"></i> Copy to Clipboard
                </span>
            </div>
        </div>
        <pre ref="text">
| **Q** | **A** |
|---|---|
| JIRA | {{jiraFormatted}} |
| Linked PR | {{linkedPrFormatted}} |
| Related PR | {{relatedPrFormatted}} |
| CMS | {{cmsText}} |
| DB | {{dbText}} |
| Live Follow-up | {{liveFollowUpText}} |
| Parameters | {{parametersText}} |
| Requires cloudinary upload | {{cloudinaryText}} |

# Description

# How to test

# QA
        </pre>
    </div>
</template>

<script type="module">
    function jiraFormat(str) {
        str = str.trim()
        if (str === '') {
            return '_N/A_';
        }
        let strArray = str.split(',');
        for (let i = 0; i < strArray.length; i++) {
            let v = strArray[i].trim();
            strArray[i] = '**[' + v + '](https://feelunique.atlassian.net/browse/' + v + ')**';
        }
        return strArray.join(' ');
    }

    function format(str) {
        str = str.trim();
        if (str === '') {
            return '_N/A_';
        }
        let strArray = str.split(',');
        for (let i = 0; i < strArray.length; i++) {
            let v = strArray[i].trim();
            strArray[i] = '**' + v + '**';
        }
        return strArray.join(' ');
    }

    function check(b) {
        if (b) {
            return ':heavy_check_mark:';
        }
        return ':heavy_multiplication_x:';
    }

    export default {
        data() {
            return {
                jira: '',
                jiraFormatted: '_N/A_',
                linkedPr: '',
                linkedPrFormatted: '_N/A_',
                relatedPr: '',
                relatedPrFormatted: '_N/A_',
                cms: false,
                cmsText: ':heavy_multiplication_x:',
                db: false,
                dbText: ':heavy_multiplication_x:',
                liveFollowUp: false,
                liveFollowUpText: ':heavy_multiplication_x:',
                parameters: false,
                parametersText: ':heavy_multiplication_x:',
                cloudinary: false,
                cloudinaryText: ':heavy_multiplication_x:',
                clipBoardStatus: 'copy-to-clipboard-default'
            }
        },
        methods: {
            jiraChange() {
                this.jiraFormatted = jiraFormat(this.jira);
            },
            linkedPrChange() {
                this.linkedPrFormatted = format(this.linkedPr);
            },
            relatedPrPrChange() {
                this.relatedPrFormatted = format(this.relatedPr);
            },
            csmChange() {
                this.cmsText = check(this.cms);
            },
            dbChange() {
                this.dbText = check(this.db);
            },
            liveFollowUpChange() {
                this.liveFollowUpText = check(this.liveFollowUp);
            },
            parametersChange() {
                this.parametersText = check(this.parameters);
            },
            cloudinaryChange() {
                this.cloudinaryText = check(this.cloudinary);
            },
            clearField() {
                this.jira = "";
                this.jiraFormatted = "_N/A_"
                this.linkedPr = "";
                this.linkedPrFormatted = "_N/A_";
                this.relatedPr = "";
                this.relatedPrFormatted = "_N/A_";
                this.cms = false;
                this.cmsText = ':heavy_multiplication_x:';
                this.db = false;
                this.dbText = ':heavy_multiplication_x:';
                this.liveFollowUp = false;
                this.liveFollowUpText = ':heavy_multiplication_x:';
                this.parameters = false;
                this.parametersText = ':heavy_multiplication_x:';
                this.cloudinary = false;
                this.cloudinaryText = ':heavy_multiplication_x:';
            },
            clipboard() {
                navigator.clipboard.writeText(this.$refs.text.innerHTML.trim());
                this.clipBoardStatus = 'copy-to-clipboard-success';
                let copy = this;
                setTimeout(function () {
                    copy.clipBoardStatus = 'copy-to-clipboard-default';
                }, 3000);
            }
        }
    }
</script>