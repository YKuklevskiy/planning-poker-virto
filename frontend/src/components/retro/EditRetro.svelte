<script>
    import SolidButton from '../SolidButton.svelte'
    import Modal from '../Modal.svelte'
    import { _ } from '../../i18n.js'

    export let toggleEditRetro = () => {}
    export let handleRetroEdit = () => {}
    export let retroName = ''
    export let joinCode = ''

    function saveRetro(e) {
        e.preventDefault()

        const retro = {
            retroName,
            joinCode,
        }

        handleRetroEdit(retro)
    }
</script>

<Modal closeModal="{toggleEditRetro}" widthClasses="md:w-2/3 lg:w-3/5 xl:w-1/2">
    <form on:submit="{saveRetro}" name="createRetro">
        <div class="mb-4">
            <label
                class="block text-gray-700 dark:text-gray-400 text-sm font-bold mb-2"
                for="retroName"
            >
                {$_('retroName')}
            </label>
            <div class="control">
                <input
                    name="retroName"
                    bind:value="{retroName}"
                    placeholder="{$_('retroNamePlaceholder')}"
                    class="bg-gray-100 dark:bg-gray-900 border-gray-200 dark:border-gray-800 border-2 appearance-none
                rounded w-full py-2 px-3 text-gray-700 dark:text-gray-300 leading-tight
                focus:outline-none focus:bg-white dark:focus:bg-gray-700 focus:border-indigo-500 focus:caret-indigo-500 dark:focus:border-yellow-400 dark:focus:caret-yellow-400"
                    id="retroName"
                    required
                />
            </div>
        </div>

        <div class="mb-4">
            <label
                class="block text-gray-700 dark:text-gray-400 text-sm font-bold mb-2"
                for="joinCode"
            >
                {$_('passCode')}
            </label>
            <div class="control">
                <input
                    name="joinCode"
                    bind:value="{joinCode}"
                    placeholder="{$_('optionalPasscodePlaceholder')}"
                    class="bg-gray-100 dark:bg-gray-900 border-gray-200 dark:border-gray-800 border-2 appearance-none
                rounded w-full py-2 px-3 text-gray-700 dark:text-gray-300 leading-tight
                focus:outline-none focus:bg-white dark:focus:bg-gray-700 focus:border-indigo-500 focus:caret-indigo-500 dark:focus:border-yellow-400 dark:focus:caret-yellow-400"
                    id="joinCode"
                />
            </div>
        </div>

        <div class="text-right">
            <SolidButton type="submit">{$_('save')}</SolidButton>
        </div>
    </form>
</Modal>
