import Fuse from 'fuse.js'
import { useEffect, useRef, useState } from 'react'

interface FuseParams<T> {
  data: T[]
  keys?: Fuse.FuseOptionKey[]
  options?: Fuse.IFuseOptions<T> & CustomOptions
}

interface CustomOptions {
  showResultsWhenNoSearchTerm?: boolean
}

interface FuseResults<T> {
  results: Fuse.FuseResult<T>[]
  search: string
  setSearch: (search: string) => void
}

const defaultOptions = {
  shouldSort: true,
  threshold: 0.1,
  location: 0,
  distance: 100,
  maxPatternLength: 32,
  minMatchCharLength: 1,
}

const DEFAULT_QUERY = ''

export function useFuse<T>({
  data,
  keys,
  options = {},
}: FuseParams<T>): FuseResults<T> {
  const { showResultsWhenNoSearchTerm, ...fuseOptions } = options
  const [fuseResults, setFuseResults] = useState<Fuse.FuseResult<T>[]>([])
  const [search, setSearch] = useState(DEFAULT_QUERY)
  const fuse = useRef<Fuse<T>>()

  useEffect(() => {
    if (!data) return

    fuse.current = new Fuse(data, {
      ...defaultOptions,
      ...fuseOptions,
      keys,
    })
  }, [search, fuse, data, keys, options])

  useEffect(() => {
    async function set(): Promise<void> {
      if (fuse.current) {
        setFuseResults(await fuse.current.search(search))
      }
    }
    set()
  }, [search, fuse, data])

  const results =
    showResultsWhenNoSearchTerm && search === ''
      ? data.map((data, i) => ({
          item: data,
          score: 1,
          refIndex: i,
        }))
      : fuseResults

  return { results, search, setSearch }
}
