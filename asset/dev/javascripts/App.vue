<template>
    <div class="root">
        <div class="row">
            <div class="col-9">
                <div class="form-group">
                    <label for="jira">JIRA Task</label>
                    <input type="text" class="form-control form-control-sm" id="jira" aria-describedby="jiraHelp"
                           placeholder="and only the task number (including 'FU-')" v-model="jira">
                    <small id="jiraHelp" class="form-text text-muted">Seperate by comma (,)</small>
                </div>
                <div class="form-group">
                    <label for="linkedPr">Linked P.R.</label>
                    <input type="text" class="form-control form-control-sm" id="linkedPr"
                           aria-describedby="linkedPrHelp" v-model="linkedPr">
                    <small id="linkedPrHelp" class="form-text text-muted">Seperate by comma (,)</small>
                </div>
                <div class="form-group">
                    <label for="relatedPr">Related P.R.</label>
                    <input type="text" class="form-control form-control-sm" id="relatedPr"
                           aria-describedby="relatedPrHelp" v-model="relatedPr">
                    <small id="relatedPrHelp" class="form-text text-muted">Seperate by comma (,)</small>
                </div>
            </div>
            <div class="col-3">
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="CMS" v-model="cms">
                    <label class="form-check-label" for="CMS">CMS</label>
                </div>
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="DB" v-model="db">
                    <label class="form-check-label" for="DB">DB</label>
                </div>
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="LiveFollowUp" v-model="liveFollowUp">
                    <label class="form-check-label" for="LiveFollowUp">Live Follow-up</label>
                </div>
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="Parameters" v-model="parameters">
                    <label class="form-check-label" for="Parameters">Parameters</label>
                </div>
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="Cloudinary" v-model="cloudinary">
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
| JIRA | {{jira | jira}} |
| Linked PR | {{linkedPr | link}} |
| Related PR | {{relatedPr | link}} |
| CMS | {{cms | check}} |
| DB | {{db | check}} |
| Live Follow-up | {{liveFollowUp | check}} |
| Parameters | {{parameters | check}} |
| Requires cloudinary upload | {{cloudinary | check}} |

# Description

# How to test

# QA
        </pre>
    </div>
</template>

<script type="module">
    function baseFormat(str, closure) {
        str = str.trim()
        if (str === '') {
            return '_N/A_';
        }
        let strArray = str.split(',');
        for (let i = 0; i < strArray.length; i++) {
            strArray[i] = closure(strArray[i].trim());
        }
        return strArray.join(' ');
    }

    function jiraFormat(str) {
        return baseFormat(str,
            function(v) { return '**[' + v + '](https://feelunique.atlassian.net/browse/' + v + ')**'; })
    }

    function format(str) {
        return baseFormat(str, function(v) { return '**' + v + '**'; })
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
                linkedPr: '',
                relatedPr: '',
                cms: false,
                db: false,
                liveFollowUp: false,
                parameters: false,
                cloudinary: false,
                clipBoardStatus: 'copy-to-clipboard-default'
            }
        },
        methods: {
            clearField() {
                this.jira = "";
                this.linkedPr = "";
                this.relatedPr = "";
                this.cms = false;
                this.db = false;
                this.liveFollowUp = false;
                this.parameters = false;
                this.cloudinary = false;
            },
            clipboard() {
                navigator.clipboard.writeText(this.$refs.text.innerHTML.trim());
                this.clipBoardStatus = 'copy-to-clipboard-success';
                let copy = this;
                setTimeout(function () {
                    copy.clipBoardStatus = 'copy-to-clipboard-default';
                }, 3000);
            }
        },
        filters: {
            jira(value) {
                return jiraFormat(value);
            },
            link(value) {
                return format(value);
            },
            check(value) {
                return check(value);
            }
        }
    }
</script>