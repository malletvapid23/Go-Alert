import React from 'react'
import { gql } from 'urql'
import QueryList from '../lists/QueryList'
import ScheduleCreateDialog from './ScheduleCreateDialog'

const query = gql`
  query schedulesQuery($input: ScheduleSearchOptions) {
    data: schedules(input: $input) {
      nodes {
        id
        name
        description
        isFavorite
      }
      pageInfo {
        hasNextPage
        endCursor
      }
    }
  }
`

export default function ScheduleList(): JSX.Element {
  return (
    <QueryList
      query={query}
      variables={{ input: { favoritesFirst: true } }}
      mapDataNode={(n) => ({
        title: n.name,
        subText: n.description,
        url: n.id,
        isFavorite: n.isFavorite,
      })}
      CreateDialogComponent={ScheduleCreateDialog}
      createLabel='Schedule'
    />
  )
}
